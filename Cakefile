fs            = require 'fs'
path          = require 'path'
{spawn, exec} = require 'child_process'

task 'build:parser', 'rebuild the Jison parser (run build first)', ->
  require 'jison'
  parser = require('./lib/coffeephp/grammar').parser
  fs.writeFile 'lib/coffeephp/parser.js', parser.generate()

