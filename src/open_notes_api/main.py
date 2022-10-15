from fastapi import FastAPI
from typing import List
from open_notes_api import schema, db, models, controller


app = FastAPI(
    title="OpenNotes API",
    description="""API desenvolvida para o **Trabalho 1** da disciplina de sistemas distribuídos."""
)


@app.on_event("startup")
async def startup():
    async with db.engine.begin() as conn:
        await conn.run_sync(models.Base.metadata.create_all)


@app.post("/notes/", response_model=schema.Note)
async def create_note(note: schema.NoteCreate):
    """Cria uma anotação."""
    return await controller.create_note(db=db.session, note=note)


@app.get("/notes/", response_model=List[schema.Note])
async def read_notes():
    """Lista todas as anotações criadas."""
    return await controller.list_notes(db.session)


@app.get("/notes/{note_id}/", response_model=schema.Note)
async def read_note(note_id: int):
    return None


@app.patch("/notes/{note_id}/", response_model=schema.Note)
async def update_note(note_id: int, note: schema.NoteUpdate):
    return await controller.update_note(db=db.session, note_id=note_id, data=note)


@app.delete("/notes/{note_id}/")
async def delete_note(note_id: int):
    """Deleta uma anotação."""
    await controller.delete_note(db=db.session, note_id=note_id)
    return None


def run_dev():
    """Atalho para rodar o servidor durante desenvolvimento."""
    import uvicorn
    uvicorn.run("open_notes_api.main:app", log_level="info", reload=True)
