# shannon

## Overview
Program to check the Shannon entropy of a string or file.

```bash
./shannon 
  -file string
        File to calculate the Shannon entropy of
  -string string
        String to calculate the Shannon entropy of

You must provide either a string of file to check
```

## Examples
The program can take a string or file.

```bash
./shannon -string 1223334444
1.8464393446710152

./shannon -file /bin/bash
6.161764074121882
```

## References / Credit

Algorithm taken from:
- https://rosettacode.org/wiki/Entropy#Go

Threat Hunting with File Entropy
- https://practicalsecurityanalytics.com/file-entropy/

Entropy Analysis a Critical Test for Malware
- https://whiteheart0.medium.com/entropy-analysis-a-critical-test-for-malwares-69939f5b8b1#:~:text=Entropy%20analysis%20do%20not%20only,the%20machine%20learning%20av%20detection.
