#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/list.h>
#include <linux/types.h>
#include <linux/slab.h>
#include <linux/sched.h>
#include <linux/string.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>

/*void recorrerHijos(struct seq_file *sf, struct task_struct *ts, int contador)
{
    struct list_head *list;
    struct task_struct *task;
    
    list_for_each(list, &ts->children)
    {
        task = list_entry(list, struct task_struct, sibling);
        pr_info("Nombre: %s PID: %d Estado: %d\n", ts->comm, ts->pid, ts->state);
        recorrerHijos(sf, task, contador + 1);
    }
}*/


typedef struct procesos
{
    int pid;
    char nombre[150];
    int estado;
    int pidpadre;
};

void procs_info_print(void)
{
    struct task_struct *task_list;
    size_t process_counter = 0;
    int tam=0;
    for_each_process(task_list)
    {
        //pr_info("Nombre: %s \t PID: %d \t Estado: %d \t PID Padre: %d\n", task_list->comm, task_list->pid, task_list->state, task_list->real_parent->pid);
        ++process_counter;
        tam++;
    }
    printk(KERN_INFO "== Number of process: %zu\n", process_counter);
}

int init_module(void)
{
    printk(KERN_INFO "[ INIT ==\n");
    /*struct task_struct *root = current;
    while (root->pid != 1)
    {
        root = root->parent;
    }
    struct seq_file *sf;
    recorrerHijos(sf, root, 1);
    */
    procs_info_print();

    return 0;
}

void cleanup_module(void)
{
    printk(KERN_INFO "== CLEANUP ]\n");
}

MODULE_LICENSE("MIT");
