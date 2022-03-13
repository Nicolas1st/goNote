import { newNoteComponent } from "./components";

// accepts an array of note objects
// can ether add new notes to the list,
// or replace the old one with the new ones
// depending on the params
export function displayNoteComponents(replace, ...notes) {
    console.log("Display notes working");
    const notesContainer = document.querySelector(".notes-container");

    if (replace) {
        notesContainer.innerHTML = "";
    }

    // adding notes
    notes.forEach((note) => {
        console.log(note);
        notesContainer.appendChild(newNoteComponent(note.Title, note.Content, note.ID));
    });
}

export function removeNoteComponent(noteID) {
    const notesContainer = document.querySelector(".notes-container");

    const elementToBeRemoved = notesContainer.querySelector($`[data-id="${noteID}"]`)

    notesContainer.removeChild(elementToBeRemoved);
}

export function updateNoteComponent(note){
    const notesContainer = document.querySelector(".notes-container")

    const newNote = newNoteComponent(note.Title, note.Content, note.ID);
    const oldNote = notesContainer.querySelector($`[data-id="${note.ID}"]`)

    notesContainer.replaceChild(newNote, oldNote);
}
