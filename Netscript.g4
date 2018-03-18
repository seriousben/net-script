// Netscript.g4
grammar Netscript;

// Parser Rules
script
   : line+
   ;

line
   : (comment | command)
   ;

comment
   : COMMENT
   ;

command
   : method ' ' url
   ;

method
   : METHOD
   ;

url
   : TEXT
   ;

// Lexer rules
COMMENT: '#' ~[\r\n]*;
METHOD: [A-Z]+;
TEXT: ~[\n\r ]+;
TERMINATOR: [\r\n]+ -> channel(HIDDEN);
