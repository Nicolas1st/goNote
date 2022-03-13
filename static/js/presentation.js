import { newNoteComponent } from "./components";

// accepts an array of note objects
// can ether add new notes to the list,
// or replace the old one with the new ones
// depending on the params
export function displayNoteComponents(replace, ...notes) {
    const notesContainer = document.querySelector(".notes-container");

    if (replace) {
        notesContainer.innerHTML = "";
    }

    // adding notes
    notes.forEach((note) => {
        notesContainer.appendChild(newNoteComponent(note.Title, note.Content, note.ID));
    });
}

export function removeNoteComponent(note) {
    const notesContainer = document.querySelector(".notes-container");

    notesContainer.removeChild(elementToBeRemoved);
}

export function updateNoteComponent(note){
    const notesContainer = document.querySelector(".notes-container")

    const newNote = newNoteComponent(note.Title, note.Content, note.ID);
    const oldNote = notesContainer.querySelector($`[data-id="${note.ID}"]`)

    notesContainer.replaceChild(newNote, oldNote);
}
