const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

console.log("Welcome to the Workout Tracker CLI!");

function workoutTracker() {
  /*rl.question('What action would you like to perform? (add, list, quit) ', (action) => {
    if (action === "add") {
      addWorkout();
    } else if (action === "list") {
      listWorkouts();
    } else if (action === "quit" || action === "ctrl-p") {
      console.log("Exiting Workout Tracker...");
      rl.close();
    } else {
      console.log("Invalid action. Please choose 'add', 'list', or 'quit'");
      workoutTracker();
    }
  });*/
  console.log('Welcome, this is your workout tracker, please select an action :s');
}

function addWorkout() {
  rl.question('What workout did you perform? ', (workout) => {
    // console.log(`Adding ${workout} to your workout list...`);
    // code to add workout to list here
    rl.question("Enter the number of sets: ", function(sets) {
      rl.question("Enter the number of reps: ", function(reps) {
        rl.question("Enter the weight ", function(weight){
          console.log(`Exercise added: ${workout}, Sets: ${sets}, Reps: ${reps}`);
            rl.close();
        });
      });
    });
    workoutTracker();
  });
}

function listWorkouts() {
  // code to list workouts here
  workoutTracker();
}

workoutTracker();



