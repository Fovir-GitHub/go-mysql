<script lang="ts">
  import AlbumItem from "./AlbumItem.svelte";
  import type { Album } from "../types/album";

  export let albums: Album[];

  async function handleRemove(id: number) {
    const res = await fetch(
      `http://localhost:8080/albums/delete/${id}`,
      { method: "POST" },
    );
    if (!res.ok) {
      throw new Error(`HTTP error: ${res.status}`);
    }
    albums = albums ? albums.filter((album) => album.id !== id) : [];
  }
</script>

<ul>
  {#each albums as album}
    <AlbumItem {album} {handleRemove} />
  {/each}
</ul>
