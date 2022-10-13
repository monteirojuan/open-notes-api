from fastapi import FastAPI
from typing import List
import schema

app = FastAPI(
    title="OpenNotes API",
    description="""API desenvolvida para o **Trabalho 1** da disciplina de sistemas distribuídos."""
)


@app.put("/notes/", response_model=schema.Note)
async def create_note():
    return None


@app.get("/notes/", response_model=List[schema.Note])
async def read_notes():
    return None


@app.get("/notes/{note_id}/", response_model=schema.Note)
async def read_note(note_id: int):
    return None


@app.delete("/notes/{note_id}/")
async def delete_note(note_id: int):
    return None
