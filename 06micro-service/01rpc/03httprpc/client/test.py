import requests

request = {
        "id":0,
        "params":[[1,2,3,4]],
        "method":"HelloRpc.Add"
      }
rsp = requests.post("http://localhost:8000/jsonrpc", json=request)
print(rsp.text)