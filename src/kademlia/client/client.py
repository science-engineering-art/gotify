import socket

host = '18.62.255.255'
port = 8888

client_socket = socket.socket(socket.AF_INET, \
    socket.SOCK_DGRAM)

while True:
    msg = input('Enter message to send: ')
    client_socket.sendto(msg.encode('utf-8'), (host, port))
    data, addr = client_socket.recvfrom(1024)
    print(f"Received: {data}, from: {addr}")
