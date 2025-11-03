<script lang="ts">
  import { onMount } from "svelte";

  interface Album {
    id: number;
    title: string;
    artist: string;
    price: number;
  }

  let albums: Album[] | null = null;
  let error: string | null = null;

  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8080/albums/all");
      if (!res.ok) {
        throw new Error(`HTTP ${res.status}`);
      }
      albums = (await res.json()) as Album[];
    } catch (e: any) {
      error = e.message;
    }
  });

  async function handleRemove(id: number) {
    try {
      const res = await fetch(
        `http://localhost:8080/albums/delete/${id}`,
        { method: "POST" },
      );

      if (!res.ok) {
        throw new Error(`HTTP error: ${res.status}`);
      }
      albums = albums
        ? albums.filter((album) => album.id !== id)
        : null;
    } catch (e: any) {
      error = e.message;
    }
  }
</script>

<main>
  {#if error}
    <p style="color: red">{error}</p>
  {:else if !albums}
    <p>Loading albums...</p>
  {:else}
    <ul>
      {#each albums as album}
        <div>
          <li>
            <strong
              >{album.title} by {album.artist} - ${album.price}</strong
            >
          </li>
          <button onclick={() => handleRemove(album.id)}>Remove</button>
        </div>
      {/each}
    </ul>
  {/if}
</main>
