#!/usr/bin/env python

def battleship(v):
    l = 'MISS'
    if v:
        l = 'HIT'
    return '{0} : {1}'.format(l, repr(v))

if __name__ == '__main__':
    # WTF OMIT
    print '\n'.join([ battleship(x) for x in [0, '', (), [], None] ])
    # WTFE OMIT
