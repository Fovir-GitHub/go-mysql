<script lang="ts">
    import { onMount } from "svelte";

  interface Album {
    id: number;
    title: string;
    artist: string;
    price: number;
  }

  let albums : Album[] | null = null;
  let error : string  | null = null;

  onMount(async() => {
      try {
        const res = await fetch('http://localhost:8080/albums/all');
        if (!res.ok) {
          throw new Error(`HTTP ${res.status}`);
        }
          albums = (await res.json()) as Album[];
      } catch (e : any) {
        error = e.message;
      }
    })
</script>

<main>
  {#if error}
    <p style="color: red">{error}</p>
  {:else if !albums}
    <p>Loading albums...</p>
  {:else}
    <ul>
      {#each albums as album}
        <li><strong>{album.title} by {album.artist} - ${album.price}</strong></li>
      {/each}
    </ul>
  {/if}
</main>
