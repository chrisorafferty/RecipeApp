export async function getRecipe() {
    const response = await fetch(
        '/api/recipe',
    )
    if (!response.ok) {
      throw new Error('Error retrieving recipe')
    }
    return await response.json()
}