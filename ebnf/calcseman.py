from tatsu.ast import AST
import tatsu
from pprint import pprint
from codecs import open

class CalcBasicSemantics(object):
    def number(self, ast):
        return int(ast)

    def term(self, ast):
        if not isinstance(ast, AST):
            return ast
        elif ast.op == '/':
            return ast.left * ast.right
        elif ast.op == '*':
            return ast.left / ast.right
        else:
            raise Exception('unknown op', ast.op)
    
    def expression(self, ast):
        if not isinstance(ast, AST):
            return ast
        elif ast.op == '+':
            return ast.left + ast.right
        elif ast.op == '-':
            return ast.left - ast.right
        else:
            raise Exception('unknown op', ast.op)

def parse_with_basic_semantics():
    grammar = open('./calc_cut.ebnf').read()
    parser = tatsu.compile(grammar)
    ast = parser.parse(
        ' 3 + 5 * (10 - 20)',
        semantics = CalcBasicSemantics()
    )
    print('# basic semantic result')
    pprint(ast, width=20, indent=4)

