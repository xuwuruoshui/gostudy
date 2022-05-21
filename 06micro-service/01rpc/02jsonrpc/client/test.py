import json
import socket

request = {
        "id":0,
        "params":[{
            "Width":10,
            "Height":5
          }],
        "method":"Rect.Area"
      }
client = socket.create_connection(("127.0.0.1",8000))
client.sendall(json.dumps(request).encode())

rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)
