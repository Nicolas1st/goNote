import { components } from "./components";

// accepts an array of note objects
// can ether add new notes to the list,
// or replace the old one with the new ones
// depending on the params

export const presentation = {
    displayNoteComponents: function (replace, ...notes) {
        const notesContainer = document.querySelector(".notes-container");

        if (replace) {
            notesContainer.innerHTML = "";
        }

        // adding notes
        notes.forEach((note) => {
            notesContainer.appendChild(
                components.newNoteComponent(note.Title, note.Content, note.ID)
            );
        });
    },

    removeNoteComponent: function (note) {
        const notesContainer = document.querySelector(".notes-container");

        notesContainer.removeChild(note);
    },

    updateNoteComponent: function (note) {
        const notesContainer = document.querySelector(".notes-container");

        const newNote = components.newNoteComponent(note.Title, note.Content, note.ID);
        const oldNote = notesContainer.querySelector(`[data-id="${note.ID}"]`);

        notesContainer.replaceChild(newNote, oldNote);
    },
};
