import {createServer} from 'node:http';

const port = parseInt(process.env.PORT, 10) || 8080;

const server = createServer((req, res) => {
  res.end('Hello World!');
});

server.listen(port, () => {
  console.log(`Listening on ${port}`);
});
