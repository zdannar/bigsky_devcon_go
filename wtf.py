#!/usr/bin/env python

def battleship(v):
    l = 'MISS'
    if v:
        l = 'HIT'
    return '{0} : {1}'.format(l, repr(v))

if __name__ == '__main__':
    for v in [0, '', (), [], None]:
        print battleship(v)
