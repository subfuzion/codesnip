import {createServer} from 'node:http';

const port = parseInt(process.env.PORT, 10) || 8080;

const server = createServer((req, res) => {
  // Roll the dice!
  const roll = Math.floor(Math.random() * 6 + 1)

  const data = {
    roll,
    message: `🎲 You rolled a: ${roll}`,
  }

  res.writeHead(200, {'Content-Type': 'application/json'});
  res.end(JSON.stringify(data));
});

server.listen(port, () => {
  console.log(`Listening on ${port}`);
});