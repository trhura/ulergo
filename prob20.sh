#! /bin/sh

python -c 'import operator; print reduce(lambda x, y: int(x) + int(y), str(reduce(operator.mul, [x for x in range(1, 101)])))'
