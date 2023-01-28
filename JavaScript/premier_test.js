const readline = require('readline');
const fs = require('fs');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('Entrez le chemin du répertoire: ', (answer) => {
  fs.readdir(answer, (err, files) => {
    if (err) {
      console.error(err);
      return;
    }
    console.log(`Les fichiers dans ${answer} sont:`);
    console.log(files);
  });
  rl.close();
});
