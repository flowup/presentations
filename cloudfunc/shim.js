const main = spawn('./main'); // spawn the binary

main.stdin.write(JSON.stringify(body)); // write input to the stdin

let result = '';
main.stdout.on('data', (data) => {
    result += data;
});

main.on('close', (code) => { // wait for stdout and send back
    res.send(JSON.parse(result));
});