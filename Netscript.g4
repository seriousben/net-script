// Netscript.g4
grammar Netscript;

// Parser Rules
script
   : line+ EOF
   ;

line
   : (comment | command)
   ;

comment
   : COMMENT
   ;

command
   : method ' ' arguments
   ;

method
   : METHOD
   ;

arguments
   : TEXT
   ;

// Lexer rules
COMMENT: '#' ~[\r\n]*;
METHOD: [A-Z]+;
TEXT: ~[\n\r ]+;
TERMINATOR: [\r\n]+ -> channel(HIDDEN);
