/**
 * HTTP post method
 */

package xhttp

import (
    "bytes"
    "net/http"
    "io/ioutil"
)

/*
 * Post method
 * @params  postBody    body json data
 * @params  u   url
 * @params  header  HTTP header
 * @return  []byte  response data
 * @return  error   error
 */
func Post(postBody []byte, u string, header map[string] string) ([]byte, error) {
    client := &http.Client{}
    request, err := http.NewRequest("POST", u, bytes.NewBuffer(postBody))

    for key, value := range header {
        request.Header.Set(key, value)
    }

    response, err := client.Do(request)
    defer response.Body.Close()

    if err != nil {
        return []byte(""), err
    }

    data, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return []byte(""), err
    }

    return data, nil
}
