/**
 * HTTP get method
 */

package xhttp

import (
    "net/http"
    "io/ioutil"
)

/*
 * Get method
 * @params  u   url
 * @return  []byte  response data
 * @return  error   error
 */
func Get(u string) ([]byte, error) {
    response, err := http.Get(u)

    if err != nil {
        return []byte(""), err
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return []byte(""), err
    }

    return body, nil
}
