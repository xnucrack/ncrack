#import "library.h"

@implementation Cracker

- (void)helloThere {
    NSLog(@"Hello from dylib");
}

- (void)toReplace {
    NSLog(@"I am not replaced");
}

@end
