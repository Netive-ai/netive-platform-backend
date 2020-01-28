const express    =    require('express');
const app        =    express();
const db         =    require('./appDatabase');
const router     =    require('./appRoutes');


async function main()
{
    await db.connect();
    await db.seed();

    app.use(router);
    app.use(express.static(__dirname + '/public'));
    app.listen(3000, () => console.log("API is listening."));
}

main()
module.exports = app