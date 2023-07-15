
const tasks = [];
const SLAVE_COUNT = 4;

const slaves = {};


// Create slave workers
for(let i = 0; i < SLAVE_COUNT; i++) {
    const slave = new Worker('slave.js');
    slaves[`slave${i+1}`] = {
        worker: slave,
        task: null,
        name: `slave${i+1}`
    }

    slave.onmessage = function(e) {
        postMessage(e.data);
        delete e;
        slaves[`slave${i+1}`].task = null;     
        manager();
    }
}

function getNextSlave() {
    for(let key in slaves) {
        if(slaves[key].task === null) {
            return slaves[key];
        }
    }
    return null;
}

function manager(){
    const nextSlave = getNextSlave();
    if(nextSlave === null) {
        console.log('no slave available');
    }
    else {
        const task = tasks.shift();
        if(task) {
            nextSlave.task = task;
            // send task to slave
            nextSlave.worker.postMessage({image:task, name: nextSlave.name});
        }
        delete task;
    }
}


// Geg the message from webpage
onmessage = function(e) {
    console.log(e);
    if(e.data.from === 'webpage') {
        tasks.push(e.data.imagebase64);
        manager();
    }
}



