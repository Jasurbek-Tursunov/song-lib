# Song Library API

## Overview
The **Song Library API** is a RESTful service designed for managing and retrieving information about songs, including their metadata and lyrics. It provides endpoints for CRUD operations on songs, as well as pagination for large datasets.

## API Version
- Version: `1.0`

## Run project
- Make sure there is a .env file. Create it from the example example.env file or run command 
```bash
  make copy-default-env
```
---

- Use them to run in developer mode on different terminals
```bash
    make run-mock-external
```

```bash
    make run
```
---
- Use for release
```bash
    make build-release
    make run-release
```


## Endpoints

### 1. **List Songs**
- **Endpoint:** `GET /songs`
- **Description:** Retrieves a list of songs with optional filters and pagination.
- **Parameters:**
    - **Query Parameters:**
        - `group` *(string)*: Filter by group name.
        - `song` *(string)*: Filter by song name.
        - `release-date` *(string)*: Filter by release date.
        - `link` *(string)*: Filter by song link.
        - `limit` *(integer)*: Number of items per page.
        - `page` *(integer)*: Page number for pagination.
- **Responses:**
    - `200 OK`: Returns a list of songs.
    - `400 Bad Request`: Invalid query parameters.
    - `500 Internal Server Error`: Server-side error.

---

### 2. **Create Song**
- **Endpoint:** `POST /songs`
- **Description:** Creates a new song record.
- **Body:**
    - A JSON object representing the song:
      ```json
      {
        "group": "string",
        "song": "string"
      }
      ```
- **Responses:**
    - `200 OK`: Returns the created song.
    - `400 Bad Request`: Invalid input data.
    - `404 Not Found`: Related resource not found.
    - `500 Internal Server Error`: Server-side error.

---

### 3. **Get Song**
- **Endpoint:** `GET /songs/{id}`
- **Description:** Retrieves information about a specific song by its ID.
- **Path Parameters:**
    - `id` *(integer)*: The unique identifier of the song.
- **Responses:**
    - `200 OK`: Returns the song details.
    - `404 Not Found`: Song not found.
    - `500 Internal Server Error`: Server-side error.

---

### 4. **Update Song**
- **Endpoint:** `PUT /songs/{id}`
- **Description:** Updates a specific song by its ID.
- **Path Parameters:**
    - `id` *(integer)*: The unique identifier of the song.
- **Body:**
    - A JSON object with the song's updated details:
      ```json
      {
        "group": "string",
        "song": "string",
        "releaseDate": "string",
        "link": "string"
      }
      ```
- **Responses:**
    - `200 OK`: Returns the updated song.
    - `400 Bad Request`: Invalid input data.
    - `404 Not Found`: Song not found.
    - `500 Internal Server Error`: Server-side error.

---

### 5. **Delete Song**
- **Endpoint:** `DELETE /songs/{id}`
- **Description:** Deletes a specific song by its ID.
- **Path Parameters:**
    - `id` *(integer)*: The unique identifier of the song.
- **Responses:**
    - `200 OK`: Song successfully deleted.
    - `404 Not Found`: Song not found.
    - `500 Internal Server Error`: Server-side error.

---

### 6. **Get Song Verses**
- **Endpoint:** `GET /songs/{id}/text`
- **Description:** Retrieves the verses (lyrics) of a specific song by its ID, with pagination.
- **Path Parameters:**
    - `id` *(integer)*: The unique identifier of the song.
- **Query Parameters:**
    - `limit` *(integer)*: Number of items per page.
    - `page` *(integer)*: Page number for pagination.
- **Responses:**
    - `200 OK`: Returns a list of song verses.
    - `400 Bad Request`: Invalid query parameters.
    - `404 Not Found`: Song not found.
    - `500 Internal Server Error`: Server-side error.

---
