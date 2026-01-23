from pathlib import Path

import dotenv

# Path to the .env in the parent 'control' folder
ENV_PATH = Path(__file__).parent.parent / ".env"


def get_variables():
    """
    Standard Robot Framework variable loader.
    Uses dotenv_values to ensure $VAR interpolation is resolved.
    """
    # dotenv_values returns a dict with all variables already interpolated
    return dict(dotenv.dotenv_values(dotenv_path=ENV_PATH, interpolate=True))
