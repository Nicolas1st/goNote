/* all fucntions defined in this file work in the following way */

/* they make requests to the api */
/* if succesfull they return a json */
/* otherwise they error out */

/* it's the calling code responsibility to check for errors */

export async function loadNotes(authorName) {
    const response = await fetch(`/notes/${authorName}/`);

    return JSON.parse(JSON.stringify(await response.json()));
}

export async function createNote(authorName, noteTitle, noteContent) {
    const response = await fetch(`/notes/`, {
        method: "POST",
        header: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Title: noteTitle,
            Content: noteContent,
            Author: authorName
        }),
    });

    return response.json();
}

export async function deleteNote(authorName, noteID) {
    const response = await fetch(`/notes/${noteID}/`, {
        method: "DELETE",
    });

    return await response.json();
}

export async function updateNote(authorName, noteID) {
    const response = await fetch(`/notes/${noteID}/`, {
        method: "PUT",
        body: JSON.stringify({
            Title: noteTitle,
            Content: noteContent,
        }),
    });

    return await response.json();
}