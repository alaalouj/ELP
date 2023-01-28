#!/usr/bin/env node

const app = require('commander');

app
.version('1.0.0')
.option('-e, --environment <environment>', 'Environment to work with')
.action(options => {
    console.log(options.environment);
});

app.parse(process.argv);