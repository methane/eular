#!/usr/bin/env rdmd
import std.stdio;
import std.functional;

int[] primes;

void calc_primes(int n) {
    primes = [2];
    for (int i=3; i<=n; i+=2) {
        foreach (int p; primes) {
            if (p*p > i) {
                primes ~= i;
                break;
            }
            if (i % p == 0) {
                break;
            }
        }
    }
}

/*
int d(int n) {
    writeln("d(", n, ")");
    int[] divisor_primes;
    foreach (int p; primes) {
        if (n < 2) break;
        while (n % p == 0) {
            divisor_primes ~= p;
            n /= p;
        }
    }
    int sum=0;  // 1 is one of divisors.
    for (int i=0; i< (1 << divisor_primes.length); ++i) {
        int prod=1;
        for (int j=0; j < divisor_primes.length; ++j) {
            if (((1 << j) & i) != 0) {
                prod *= divisor_primes[j];
            }
        }
        write(prod, " ");
        sum += prod;
    }
    writeln("\n=> ", sum);
    return sum;
}
*/

int d(int n) {
    int sum = 0;
    for (int i = 1; i < n; ++i) {
        if ((n % i) == 0) {
            sum += i;
        }
    }
    //writefln("d(%d) == %d", n, sum);
    return sum;
}

alias memoize!(d, 10000) md;

bool is_amicable(int n) {
    int mn = md(n);
    if (mn == n) return false;
    return n == md(mn);
}

void main() {
    long sum = 0;
    for (int i=1; i<10000; ++i) {
        if (is_amicable(i)) {
            writeln(i);
            sum += i;
        }
    }
    writeln(sum);
}
