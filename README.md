# Family Tree Command Line Tool Documentation

## Introduction

The Family Tree Command Line Tool is a Go-based application designed to manage and interact with a family tree through a command line interface. This tool allows users to add individuals, establish relationships between them, and query information about the family structure. The family tree data is stored in a YAML file for persistent storage.

## Getting Started

### Installation

To install the Family Tree Command Line Tool, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/VILJkid/go-family-tree.git
   ```

2. Change into the project directory:

   ```bash
   cd go-family-tree
   ```

3. Build the executable:

   ```bash
   go build
   ```

4. Run the executable:

   ```bash
   ./go-family-tree [command] [arguments]
   ```

### Command Structure

The tool supports the following commands:

- **add**: Add a new person or relationship.
- **connect**: Establish a relationship between two individuals.
- **count**: Retrieve the count of specific relationships for a person.
- **find**: Find individuals related to a person based on a specific relationship.

## Commands

### 1. Add Command

#### 1.1 Add Person

```bash
./go-family-tree add person [Name]
```

- Adds a new person to the family tree with the specified name.

#### 1.2 Add Relationship

```bash
./go-family-tree add relationship
```

- This functionality is built-in and not directly invoked by the user.

### 2. Connect Command

```bash
./go-family-tree connect [Person1] as [Relationship] of [Person2]
```

- Establishes a relationship between two individuals in the family tree.

#### Examples:

- Connect John as son of Mary:

  ```bash
  ./go-family-tree connect John as son of Mary
  ```

- Connect Jane as daughter of John:

  ```bash
  ./go-family-tree connect Jane as daughter of John
  ```

### 3. Count Command

```bash
./go-family-tree count [Relationship] of [Person]
```

- Retrieves the count of a specific relationship for a given person.

#### Examples:

- Count the number of sons for John:

  ```bash
  ./go-family-tree count sons of John
  ```

- Count the number of wives for Mary:

  ```bash
  ./go-family-tree count wives of Mary
  ```

### 4. Find Command

```bash
./go-family-tree find [Relationship] of [Person]
```

- Finds individuals related to a person based on a specific relationship.

#### Examples:

- Find the sons of John:

  ```bash
  ./go-family-tree find sons of John
  ```

- Find the daughters of Mary:

  ```bash
  ./go-family-tree find daughters of Mary
  ```

## File Structure

- **family/types.go**: Defines the data structures used for representing individuals and relationships in the family tree.

- **family/constants.go**: Contains constant values such as file names and command keywords.

- **family/family.go**: Implements the core functionality of the family tree command line tool, including parsing, adding, connecting, counting, and finding operations.

- **main.go**: Contains the main entry point for the command line tool, invoking the appropriate command based on user input.

## Data Storage

The family tree data is stored in a YAML file named `people.yml`. This file is created and updated as individuals and relationships are added or modified using the command line tool.

## Conclusion

The Family Tree Command Line Tool provides a simple and efficient way to manage and explore family relationships. Users can add individuals, establish connections, and query information about their family structure, all from the command line.
