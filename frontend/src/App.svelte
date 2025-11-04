<script lang="ts">
  import { onMount } from "svelte";
  import AlbumList from "./components/AlbumList.svelte";
  import type { Album } from "./types/album";

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
</script>

<main>
  {#if error}
    <p style="color: red">{error}</p>
  {:else if !albums}
    <p>Loading albums...</p>
  {:else}
    <AlbumList bind:albums />
  {/if}
</main>
