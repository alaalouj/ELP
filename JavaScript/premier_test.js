const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

let exercises = [];

function addExercise() {
  rl.question('Enter the name of the exercise: ', (exercise) => {
    exercises.push(exercise);
    console.log(`${exercise} has been added to your workout.`);
    mainMenu();
  });
}

function viewExercises() {
  console.log(`Your exercises: ${exercises.join(', ')}`);
  mainMenu();
}

function deleteExercise() {
  rl.question('Enter the name of the exercise to delete: ', (exercise) => {
    const index = exercises.indexOf(exercise);
    if (index !== -1) {
      exercises.splice(index, 1);
      console.log(`${exercise} has been deleted from your workout.`);
    } else {
      console.log(`${exercise} was not found in your workout.`);
    }
    mainMenu();
  });
}

function mainMenu() {
  rl.question('What would you like to do? (Add exercise, View exercises, Delete exercise, Exit)', (answer) => {
    switch (answer) {
      case 'Add exercise':
        addExercise();
        break;
      case 'View exercises':
        viewExercises();
        break;
      case 'Delete exercise':
        deleteExercise();
        break;
      case 'Exit':
        rl.close();
        break;
      default:
        console.log('Invalid choice. Please choose again.');
        mainMenu();
    }
  });
}

console.log('Welcome to the Workout Tracker CLI!');
mainMenu();
