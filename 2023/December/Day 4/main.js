const args = process.argv.slice(2);
const isTest = args.includes('--test');
const fs = require('fs');