import { useQuery } from "@tanstack/react-query";
import { getRecipe } from "./api/recipe";

function Recipe() {
    const { isPending, isError, error, data } = useQuery({
        queryKey: ['repoData'],
        queryFn: getRecipe,
    });

    if (isPending) {
        return <span>Loading...</span>;
    }

    if (isError) {
        return <span>Error: {error.message}</span>;
    }

    return <div>{JSON.stringify(data)}</div>;
}



export default Recipe;