#!/usr/bin/env bash
# empty diff is ok
diff <(../generator/generator -n 149876|sort) <(../generator/generator -n 149876|./sort)
