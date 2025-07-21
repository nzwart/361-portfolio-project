<script>
    function openBattleManager() {
        window.location.href = "/encbuilder";
    }
</script>

<div class="container">
    <h1 class="frontpage-title">Battle Manager</h1>

    <div class="body">
        <p>
            Battle Manager is a webapp for building and tracking 5th Edition
            Dungeons & Dragons combat encounters on the fly.
        </p>
        <p>
            Search all monsters freely available under the Open Gaming License
            and add them to your encounter.
        </p>
        <p>Track their health and initiative for easy combat management!</p>
    </div>

    <div class="button-container">
        <button on:click={openBattleManager}>Open Battle Manager</button>
    </div>
</div>
