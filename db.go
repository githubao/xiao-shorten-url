// storage
// author: baoqiang
// time: 2019-06-19 19:10
package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"goji.io/pat"
	"net/http"
	"net/url"
)

type DBHandler struct {
	DB *bolt.DB
}

func (h *DBHandler) getShortUrl(w http.ResponseWriter, r *http.Request) {
	shortURL := pat.Param(r, "surl")

	if len(shortURL) != 6 {
		fmt.Fprint(w, "404: not found !")
		return
	}

	var longURL string

	_ = h.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("shortURL"))
		v := b.Get([]byte(shortURL))
		if len(v) > 0 {
			longURL = string(v)
		}

		return nil
	})

	if len(longURL) > 0 {
		http.Redirect(w, r, longURL, http.StatusSeeOther)
		return
	}

	fmt.Fprint(w, "404: not found !")
}

func (h *DBHandler) api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=utf-8")

	err := r.ParseForm()
	if err != nil {
		responseErr(ErrnoParamIllegal, w)
		return
	}

	longURL := r.FormValue("long_url")
	if len(longURL) == 0 {
		responseErr(ErrnoParamIllegal, w)
		return
	}

	_, err = url.ParseRequestURI(longURL)
	if err != nil {
		responseErr(ErrnoParamUrlFailed, w)
		return
	}

	urlList := URLShorten(longURL)

	curSURL := ""
	surlExist := false //longUrl exists
	hasSURL := false   // 有短url可用
	surlMap := make(map[string]int)

	if len(urlList) > 0 {
		_ = h.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("shortURL"))

			for _, surl := range urlList {
				curSURL = surl
				v := b.Get([]byte(surl))

				vs := string(v)
				if len(v) > 0 {
					if vs == longURL {
						// exists
						surlMap[surl] = 2
						surlExist = true
						break
					} else {
						// conflict
						surlMap[surl] = 1
					}
				} else {
					// hint one
					surlMap[surl] = 0
					hasSURL = true
					break
				}

			}

			return nil
		})
	}

	if surlExist {
		result := fmt.Sprintf("http://%s/%s", r.Host, curSURL)
		responseSuccess(OK, result, w)
		return
	}

	// add new one
	if hasSURL {
		_ = h.DB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("shortURL"))
			for k,v := range surlMap{
				if v == 0{
					_ = b.Put([]byte(k), []byte(longURL))
					result := fmt.Sprintf("http://%s/%s", r.Host, k)
					responseSuccess(OK,result,w)
					break
				}
			}

			return nil
		})
	}else {
		responseErr(ErrnoParamNoAvailable,w)
	}

}

func responseErr(code int64, w http.ResponseWriter) {
	resp := Response{
		ErrCode: code,
		ErrMsg:  getErrMsg(code),
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func responseSuccess(code int64, result string, w http.ResponseWriter) {
	resp := Response{
		ErrCode: code,
		ErrMsg:  getErrMsg(code),
		Result:  result,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

