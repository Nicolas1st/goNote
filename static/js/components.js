export function newNoteComponent(noteTitle, noteContent, noteID) {
    // note itself
    const note = document.createElement("div");
    note.classList.add("note");

    // note's title
    const title = document.createElement("h2");
    title.innerText = noteTitle;
    title.classList.add("title");

    // note's content
    const content = document.createElement("p");
    content.innerText = noteContent;
    content.classList.add("content");

    // assembling
    note.appendChild(title);
    note.appendChild(content);

    // setting data attributes
    note.dataset.id = noteID; 

    return note;
}
