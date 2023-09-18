/**
 * @Author: anchnet
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/7/2 13:07
 */

package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"net"
	"strconv"
	"time"

	"net/http"
)

func main() {
	start := time.Now() // 获取当前时间

	//支持参数
	var (
		count int    // 起始数
		total int    // 截至数
		index string // index
		id    string // id
		title string // title
	)
	flag.IntVar(&count, "c", 1, "起始数")
	flag.IntVar(&total, "e", 1, "截至数")
	flag.StringVar(&index, "i", "", "index")
	flag.StringVar(&id, "d", "", "id")
	flag.StringVar(&title, "t", "", "title")
	// 解析参数
	flag.Parse()
	if index == "" {
		index = "demo"
	}
	if id == "" {
		id = "id_1"
	}
	if title == "" {
		title = "安畅"
	}
	fmt.Println("count：", count)
	fmt.Println("total：", total)
	fmt.Println("index：", index)
	fmt.Println("id：", id)
	fmt.Println("title：", title)
	rootCAs, _ := x509.SystemCertPool()

	if rootCAs == nil {

		rootCAs = x509.NewCertPool()

	}
	addresses := []string{"https://10.233.51.63:9200"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "elastic",
		Password:  "d5LZniOld717e3U1Re7mT769",
		CloudID:   "",
		APIKey:    "",
		Transport: &http.Transport{

			MaxIdleConnsPerHost: 10,

			ResponseHeaderTimeout: time.Second,

			DialContext: (&net.Dialer{

				Timeout: 30 * time.Second,

				KeepAlive: 30 * time.Second,
			}).DialContext,

			TLSClientConfig: &tls.Config{

				InsecureSkipVerify: true, //跳过证书认证

				RootCAs: rootCAs,
			},
		},
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		fmt.Println(err, "Error creating the client")
	}

	//Get(*es, index, id)
	//Update(*es, index, id)
	//Get(*es, index, id)
	create(*es, index, count, total)
	//Search(*es, index, title)

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)

}

func create(es elasticsearch.Client, index string, count int, total int) bool {
	//var wg sync.WaitGroup
	// Create creates a new document in the index.
	// Returns a 409 response when a document with a same ID already exists in the index.
	for i := count; i < total; i++ {
		//wg.Add(1)
		k := strconv.Itoa(i)
		var buf bytes.Buffer

		content, err := ioutil.ReadFile("./test.log") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		doc := map[string]interface{}{
			"title":   "安畅是一家怎么样的公司呢？" + k,
			"content": content,
			"time":    time.Now().Unix(),
			"date":    time.Now(),
		}
		if err := json.NewEncoder(&buf).Encode(doc); err != nil {
			fmt.Println(err, "Error encoding doc")
			return false
		}

		func() {
			time.Sleep(1 * time.Millisecond)
			res, err := es.Create(index, "idx_"+k, &buf)
			if err != nil {
				fmt.Println(err, "Error create response")
			}
			//wg.Done()
			defer res.Body.Close()
			fmt.Println(res.String())
		}()
	}
	//wg.Wait()
	return true
}

//go run main.go -c 3 -e 10000
