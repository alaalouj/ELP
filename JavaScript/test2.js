const inquirer = require('inquirer');

const workoutChoices = ['Run', 'Swim', 'Bike', 'Lift'];

const mainMenu = () => {
  inquirer
    .prompt([
      {
        type: 'list',
        name: 'workout',
        message: 'What workout would you like to do today?',
        choices: workoutChoices,
      },
    ])
    .then((answers) => {
      console.log(`Great choice! Time to go ${answers.workout}!`);
    });
};

const runInfiniteLoop = () => {
    while (true) {
      mainMenu();
    }
}

runInfiniteLoop();
