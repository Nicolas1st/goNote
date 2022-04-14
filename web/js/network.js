/* all fucntions defined in this file work in the following way */

/* they make requests to the api */
/* if succesfull they return a json */
/* otherwise they error out */

/* it's the calling code responsibility to check for errors */

export const network = {
    loadNotes: async function (authorName) {
        const response = await fetch(`/notes/${authorName}/`);

        return JSON.parse(JSON.stringify(await response.json()));
    },

    createNote: async function (authorName, noteTitle, noteContent) {
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
    },

    deleteNote: async function (noteID) {
        const response = await fetch(`/notes/${noteID}/`, {
            method: "DELETE",
        });

        return await response.json();
    },

    updateNote: async function (noteID, noteTitle, noteContent) {
        const response = await fetch(`/notes/${noteID}/`, {
            method: "PUT",
            body: JSON.stringify({
                Title: noteTitle,
                Content: noteContent,
            }),
        });

        return await response.json();
    },
}