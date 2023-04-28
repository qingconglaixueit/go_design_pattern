// @Author Bing 
// @Desc
package main

type Client struct{

}

func (c *Client)InsertUsbToComputer(com Computer){
	com.InsertUsb()
}
