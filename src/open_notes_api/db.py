from sqlalchemy.ext.asyncio import async_sessionmaker, create_async_engine

engine = create_async_engine("postgresql+asyncpg://postgres:password@localhost/postgres", echo=True)
session = async_sessionmaker(autoflush=False, expire_on_commit=False)
session.configure(bind=engine)
