
export async function client(endpoint: string, body: any | null = null) {
  console.log(body)
  const resp = await fetch(`http://api.gotify.com/${endpoint}`, {
    headers: {'Content-Type': 'application/json'},
    method: body? 'POST' : 'GET',
    body: body
  }).then(res => res.json())
    .catch(e => console.log(e))
  return resp;
}

client.get = (endpoint: string) => {
  return client(endpoint);
};

client.post = (endpoint: string, body: any) => {
  return client(endpoint, body);
}