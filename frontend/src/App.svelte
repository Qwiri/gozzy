<script lang="ts">
	import Navbar from "./Components/Elements/Navbar.svelte";
	import { Router, Link, Route } from "svelte-navigator";
	import NewQuiz from "./Components/screens/NewQuiz.svelte";
	import Quiz from "./Components/screens/Quiz.svelte";

</script>

<main>
	<Navbar/>
	<Router>
		<Route path="/">
			<h1>Gozzy</h1>
		</Route>
		<Route path="/new" component={NewQuiz} />
		<Route path="/quiz/:id" component={Quiz} />
	</Router>
</main>

<style>
</style>