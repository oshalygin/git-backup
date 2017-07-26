## 1.1.0 (July 25, 2017)

- Addition of Git Service tests

## 1.0.1 (July 21, 2017)

- Addition logos that will be used in the README
  - Bitbucket logo
  - GitHub logo
  - Golang Logo

## 1.0.0 (July 21, 2017)

- Addition of a git push to the backup repository in BitBucket.

## 0.0.8 (July 21, 2017)

- Refactored code and added a fetch step.

## 0.0.7 (July 21, 2017)

- Addition of a git pull service that iterates through the given directories and pulls the latest master branch for each repository.

## 0.0.6 (July 17, 2017)

- Refactor directoryPath to be set from the requiredFlags check.

## 0.0.5 (July 17, 2017)

- Add directory implementation to pull down a list of directories that exist within a given path.

## 0.0.4 (July 17, 2017)

- Add go-git vendor package which will be used to initiate git commands.

## 0.0.3 (July 17, 2017)

- Add logic to set a single pathDirectory value
    
## 0.0.2 (July 15, 2017)

- Add concrete concrete tests:
    - required flags implementation to check for the path argument/flag
    
## 0.0.1 (July 15, 2017)

- Add vendor packages:
    - colors: used for coloring the output with a clean API
    - goblin: used in testing